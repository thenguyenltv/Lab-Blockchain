const { ethers } = require("hardhat");

async function main() {
  const NFTMarketplaceFactory = await ethers.getContractFactory("NFTMarketplace");
  const donations = await NFTMarketplaceFactory.deploy();
  const contractAddress = await donations.deployed();

  //await donations.wait(); // Wait for deployment to complete

  console.log("Donations contract deployed at:", donations.address);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});