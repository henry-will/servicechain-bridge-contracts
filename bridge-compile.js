#!/usr/bin/env node

const fs = require("fs")
const path = require("path")

const {execFile} = require('child_process');
const util = require("util");

const TruffleConfig = require("@truffle/config");
const Contracts = require("@truffle/workflow-compile");
const AbiGen = require("@chainsafe/truffle-plugin-abigen");

const truffleConfig = TruffleConfig.detect({}, "./truffle-config.js");
console.log(`config ${JSON.stringify(truffleConfig.compilers.solc)}`);

const GENERATE = [
  {
    pkg: "bridge",
    contract: "./contracts/sc/bridge",
    build: "./build/contracts",
  },
];

(async () => {
  try {
    for (const gen of GENERATE) {
      // expected config object
      const config = {
        contracts_directory: gen.contract, // dir where contracts are located
        contracts_build_directory: gen.build // dir where contract artifacts will be saved
      };

      fs.rmdirSync("build", {recursive: true});
      fs.rmdirSync("abigenBindings", {recursive: true});
      fs.rmdirSync(gen.pkg, {recursive: true});
      fs.mkdirSync(gen.pkg, {recursive: true});

      // await execFile("rm", ["-rf", "build"]);
      // await execFile("rm", ["-rf", "abigenBindings"]);
      // await execFile("rm", ["-rf", gen.pkg]);
      //
      // await execFile("mkdir", ["-p", gen.pkg]);

      truffleConfig.merge(config)
      await Contracts.compileAndSave(truffleConfig);
      console.log("Compilation complete!");

      try {
        await execFile("yarn", ["truffle", "run", "abigen"]);
      } catch (e) {
        console.error(e);
      }

      const filenames = fs.readdirSync(gen.build);

      for (const file of filenames) {
        const name = path.basename(file, path.extname(file));

        const bin = `--bin=abigenBindings/bin/${name}.bin`;
        const abi = `--abi=abigenBindings/abi/${name}.abi`;
        const combined_json = `--combined-json=${gen.build}/${name}.json`;
        const pkg = `--pkg=${gen.pkg}`;
        const type = `--type=${name}`;
        const go_file = `--out=${gen.pkg}/${name}.go`;
        const args = [abi, bin, type, pkg, go_file];
        console.log(args);
        try {
          await execFile("./bin/abigen-v1.8.4-amd64", args);
        } catch (e) {
          console.error(e);
        }
      }
    }
  } catch (e) {
    console.error(e);
  }
})();
