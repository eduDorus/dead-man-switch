
pragma solidity ^0.5.1;
contract DeadSwitch {
    address creatorOfContract;

    struct File {
        address fileOwner;
        string ipfsHash;
        string key;
        string ping;
    }

    uint numFiles;
    mapping (uint => File) public files;

    constructor() public { 
        creatorOfContract = msg.sender;
    }

    modifier onlyCreator() {
        require(msg.sender == creatorOfContract, "Access denied. Sender not creator."); 
        // If caller is not creator, it reverts 
        _;
        // Otherwise, it continues.
    }

    modifier onlyFileOwner(uint _fileId) {
        require(files[_fileId].fileOwner == msg.sender, "Access denied. Sender not owner.");
        // If caller is not fileOwner, it reverts
        _;
        // Otherwise, it continues
    }

    function addFile(string memory _ipfsHash, address _creator) public onlyCreator returns (uint fileId) {
        fileId = numFiles++; 
        files[fileId] = File(_creator, _ipfsHash, "", "");
    }

    function pubishKey(string memory _key, uint _fileID) public onlyCreator{
        files[_fileID].key = _key;
    }

    function ping(string memory _ping, uint _fileID) public onlyFileOwner(_fileID) {
        files[_fileID].ping = _ping;
    }
}