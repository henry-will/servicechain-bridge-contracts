module.exports = {

    plugins: [
        "@chainsafe/truffle-plugin-abigen"
    ],

    compilers: {
        solc: {
            version: "0.8.14",      // Fetch exact version from solc-bin (default: truffle's version)
            // docker: true,        // Use "0.5.1" you've installed locally with docker (default: false)
        }
    },
};
