#!/usr/bin/env node

const util = require("util");
const fs = require("fs").promises;
const path = require("path");
const execFile = util.promisify(require("child_process").execFile);

const GENERATE = [
    {
        pkg: "kip13",
        directory: "kips/kip13",
    },
    {
        pkg: "bridge",
        directory: "sc/bridge",
    },
    {
        pkg: "bridge",
        directory: "sc/bridge/interface",
    },
    {
        pkg: "extbridge",
        directory: "sc/bridge/extend",
    },
    {
        pkg: "scnft",
        directory: "sc/tokens/erc721",
    },
    {
        pkg: "sctoken",
        directory: "sc/tokens/erc20",
    },
];

const GEN_FILE_DIR = "go";

(async () => {
    await execFile("rm", ["-rf", GEN_FILE_DIR]);
    await execFile("mkdir", ["-p", GEN_FILE_DIR]);

    for (const gen of GENERATE) {
        await execFile("rm", ["-rf", "build"]);
        await execFile("rm", ["-rf", "contracts"]);
        await execFile("rm", ["-rf", "abigenBindings"]);
        await execFile("mkdir", ["-p", `${GEN_FILE_DIR}/${gen.pkg}`]);
        await execFile("mkdir", ["-p", `contracts`]);
        await execFile("cp", ["-r", `contracts-src/${gen.directory}`, `contracts/${gen.pkg}`]);

        await execFile("yarn", ["truffle", "compile"]);
        await execFile("yarn", ["truffle", "run", "abigen"]);

        const files = await fs.readdir("./abigenBindings/abi/");
        for (const file of files) {
            const filename = path.basename(file, path.extname(file))
            const bin = `--bin=abigenBindings/bin/${filename}.bin`;
            const abi = `--abi=abigenBindings/abi/${filename}.abi`;
            const type = `--type=${filename}`;

            const pkg = `--pkg=${gen.pkg}`;
            const go_code = `--out=${GEN_FILE_DIR}/${gen.pkg}/${filename}.go`;
            const args = [bin, abi, type, pkg, go_code];
            console.log(args);
            await execFile("./bin/abigen-v1.9.0", args);
        }
    }
})();
