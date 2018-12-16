//Write your own contracts here. Currently compiles using solc v0.4.15+commit.bbb8e64f.
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
        require(msg.sender == creatorOfContract); // If caller is not creator, it reverts 
        _;                              // Otherwise, it continues.
    }

    modifier onlyFileOwner(uint _fileId) {
        require(files[_fileId].fileOwner == msg.sender); // If caller is not creator of specified file, it reverts
        _;                                             // Otherwise, it continues.
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