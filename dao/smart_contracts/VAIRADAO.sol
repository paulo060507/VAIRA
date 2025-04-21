// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract VAIRADAO {
    struct Proposal {
        string description;
        uint voteCount;
    }

    address public owner;
    Proposal[] public proposals;
    mapping(address => bool) public voters;

    constructor() {
        owner = msg.sender;
    }

    function createProposal(string memory _description) public {
        require(msg.sender == owner, "Apenas o criador pode propor");
        proposals.push(Proposal(_description, 0));
    }

    function vote(uint proposalIndex) public {
        require(!voters[msg.sender], "Ja votou");
        voters[msg.sender] = true;
        proposals[proposalIndex].voteCount++;
    }

    function getProposalCount() public view returns (uint) {
        return proposals.length;
    }
}
