require("@nomiclabs/hardhat-waffle");
require('dotenv').config();
const { INFURA_API_KEY, SEPOLIA_PRIVATE_KEY } = process.env;

module.exports = {
  solidity: "0.8.20",
  networks: {
    sepolia: {
      url: `https://sepolia.infura.io/v3/${INFURA_API_KEY}`,
      accounts: [SEPOLIA_PRIVATE_KEY],
    },
  },
};