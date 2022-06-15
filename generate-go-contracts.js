#!/usr/bin/env node

const util = require("util");
const execFile = util.promisify(require("child_process").execFile);

const GENERATE = [
    {
        name: "MyNFT",
        pkg: "mynfg",
        go_file: "mynfg.go",
    },
];

const GEN_FILE_DIR = `go-contracts`;

(async () => {
    await execFile("rm", ["-rf", GEN_FILE_DIR]);
    await execFile("mkdir", ["-p", GEN_FILE_DIR]);

    for (const gen of GENERATE) {
        await execFile("mkdir", ["-p", `${GEN_FILE_DIR}/${gen.pkg}`]);
        const bin = `--bin=abigenBindings/bin/${gen.name}.bin`;
        const abi = `--abi=abigenBindings/abi/${gen.name}.abi`;
        const pkg = `--pkg=${gen.pkg}`;
        const go_code = `--out=${GEN_FILE_DIR}/${gen.pkg}/${gen.go_file}`;
        const args = [bin, abi, pkg, go_code];
        console.log(args);
        await execFile("abigen", args);
    }
})();
