#!/usr/bin/env node

const util = require("util");
const execFile = util.promisify(require("child_process").execFile);

const GENERATE = [
    {
        name: "InterfaceIdentifier",
        pkg: "InterfaceIdentifier",
        directory: "contracts/kips/kip13",
        go_file: "interface_identifier.go",
    },
    {
        name: "Bridge",
        pkg: "bridge",
        directory: "contracts/sc/bridge",
        go_file: "bridge.go",
    },
    {
        name: "ExtBridge",
        pkg: "extbridge",
        directory: "contracts/sc/bridge/extend",
        go_file: "ext_bridge.go",
    },
    {
        name: "ServiceChainNFT",
        pkg: "scnft",
        directory: "contracts/sc/tokens/erc721",
        go_file: "servicechain_nft.go",
    },
    {
        name: "ServiceChainToken",
        pkg: "sctoken",
        directory: "contracts/sc/tokens/erc20",
        go_file: "servicechain_token.go",
    },
];

const GEN_FILE_DIR = "go";

(async () => {
    await execFile("rm", ["-rf", GEN_FILE_DIR]);
    await execFile("mkdir", ["-p", GEN_FILE_DIR]);

    for (const gen of GENERATE) {
        await execFile("mkdir", ["-p", `${GEN_FILE_DIR}/${gen.directory}`]);
        const bin = `--bin=abigenBindings/bin/${gen.name}.bin`;
        const abi = `--abi=abigenBindings/abi/${gen.name}.abi`;
        const combined_json = `--abi=build/contracts/${gen.name}.json`;
        const type = `--type=${gen.name}`;

        const pkg = `--pkg=${gen.pkg}`;
        const go_code = `--out=${GEN_FILE_DIR}/${gen.directory}/${gen.go_file}`;
        const args = [combined_json, type, bin, abi, pkg, go_code];
        console.log(args);
        await execFile("./bin/abigen-v1.9.0", args);
    }
})();
