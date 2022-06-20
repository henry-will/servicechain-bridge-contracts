#!/usr/bin/env node

const fs = require("fs");
const {execFile} = require('child_process');

const BINDING_PKG = [
    {
        name: "InterfaceIdentifier",
        pkg: "kip13",
        directory: "./contracts/kips/kip13",
    },
    {
        name: "Bridge",
        pkg: "bridge",
        directory: "./contracts/sc/bridge",
    },
    {
        name: "ExtBridge",
        pkg: "extbridge",
        directory: "./contracts/sc/bridge/extend",
    },
    {
        name: "ServiceChainNFT",
        pkg: "scnft",
        directory: "./contracts/sc/tokens/erc721",
    },
    {
        name: "ServiceChainToken",
        pkg: "sctoken",
        directory: "./contracts/sc/tokens/erc20",
    },
];

(async () => {
    for (const bindingPKG of BINDING_PKG) {
        fs.rmdirSync(bindingPKG.pkg, {recursive: true});
        fs.mkdirSync(bindingPKG.pkg, {recursive: true});

        const sol = `--sol=${bindingPKG.directory}/${bindingPKG.name}.sol`;

        const pkg = `--pkg=${bindingPKG.pkg}`;
        const go_file = `--out=${bindingPKG.pkg}/${bindingPKG.name}.go`;
        const args = [sol, pkg, go_file];
        console.log(args);
        await execFile("./bin/abigen-v1.9.0", args);
    }
})();
