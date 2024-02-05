// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract NFTMarketplace is ERC721 {

    uint256 public nextTokenId;
    uint256 public listingPrice = 0.000001 ether;

    constructor() ERC721("YourNFT", "YNFT") {}

    function mint() external payable {
        
        require(msg.value >= listingPrice, "Insufficient funds to mint");

        _safeMint(msg.sender, nextTokenId);
        nextTokenId++;
    }

    // function list(uint256 tokenId) external onlyOwner {
    //     require(_exists(tokenId), "Token does not exist");
    //     // Additional logic for listing, e.g., setting a flag or updating a mapping
    // }

    // function buy(uint256 tokenId) external payable {
    //     require(_exists(tokenId), "Token does not exist");
    //     require(msg.value >= listingPrice, "Insufficient funds to buy");

    //     address owner = ownerOf(tokenId);
    //     payable(owner).transfer(msg.value);
    //     _transfer(owner, msg.sender, tokenId);
    // }
}
