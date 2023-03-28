// Sources flattened with hardhat v2.13.0 https://hardhat.org

// File @openzeppelin/contracts/utils/Context.sol@v4.8.2

// SPDX-License-Identifier: MIT
// OpenZeppelin Contracts v4.4.1 (utils/Context.sol)

pragma solidity ^0.8.0;

/**
 * @dev Provides information about the current execution context, including the
 * sender of the transaction and its data. While these are generally available
 * via msg.sender and msg.data, they should not be accessed in such a direct
 * manner, since when dealing with meta-transactions the account sending and
 * paying for execution may not be the actual sender (as far as an application
 * is concerned).
 *
 * This contract is only required for intermediate, library-like contracts.
 */
abstract contract Context {
    function _msgSender() internal view virtual returns (address) {
        return msg.sender;
    }

    function _msgData() internal view virtual returns (bytes calldata) {
        return msg.data;
    }
}


// File @openzeppelin/contracts/access/Ownable.sol@v4.8.2


// OpenZeppelin Contracts (last updated v4.7.0) (access/Ownable.sol)

pragma solidity ^0.8.0;

/**
 * @dev Contract module which provides a basic access control mechanism, where
 * there is an account (an owner) that can be granted exclusive access to
 * specific functions.
 *
 * By default, the owner account will be the one that deploys the contract. This
 * can later be changed with {transferOwnership}.
 *
 * This module is used through inheritance. It will make available the modifier
 * `onlyOwner`, which can be applied to your functions to restrict their use to
 * the owner.
 */
abstract contract Ownable is Context {
    address private _owner;

    event OwnershipTransferred(address indexed previousOwner, address indexed newOwner);

    /**
     * @dev Initializes the contract setting the deployer as the initial owner.
     */
    constructor() {
        _transferOwnership(_msgSender());
    }

    /**
     * @dev Throws if called by any account other than the owner.
     */
    modifier onlyOwner() {
        _checkOwner();
        _;
    }

    /**
     * @dev Returns the address of the current owner.
     */
    function owner() public view virtual returns (address) {
        return _owner;
    }

    /**
     * @dev Throws if the sender is not the owner.
     */
    function _checkOwner() internal view virtual {
        require(owner() == _msgSender(), "Ownable: caller is not the owner");
    }

    /**
     * @dev Leaves the contract without owner. It will not be possible to call
     * `onlyOwner` functions anymore. Can only be called by the current owner.
     *
     * NOTE: Renouncing ownership will leave the contract without an owner,
     * thereby removing any functionality that is only available to the owner.
     */
    function renounceOwnership() public virtual onlyOwner {
        _transferOwnership(address(0));
    }

    /**
     * @dev Transfers ownership of the contract to a new account (`newOwner`).
     * Can only be called by the current owner.
     */
    function transferOwnership(address newOwner) public virtual onlyOwner {
        require(newOwner != address(0), "Ownable: new owner is the zero address");
        _transferOwnership(newOwner);
    }

    /**
     * @dev Transfers ownership of the contract to a new account (`newOwner`).
     * Internal function without access restriction.
     */
    function _transferOwnership(address newOwner) internal virtual {
        address oldOwner = _owner;
        _owner = newOwner;
        emit OwnershipTransferred(oldOwner, newOwner);
    }
}


// File contracts/Accounts.sol


pragma solidity 0.8.9;

pragma solidity 0.8.9;

contract Accounts is Ownable {
    string public userNmae;


    constructor(string memory _name){
        userNmae = _name;
    }

    function callContract(address contractAddress, bytes memory data) external onlyOwner returns (bytes memory) {
        (bool success, bytes memory result) = contractAddress.call(data);
        require(success, "Method call failed");
        return result;
    }
}

contract AccountsFactory {
    /**
   * @dev Deploy the contract with the desired content
    * @param _opcode Only opcode
    * @param _name Contract name
    */
    function createPair(
        uint256 _opcode,
        string memory _name
    ) public returns (address pair) {
        (bytes memory bytecode, bytes32 salt) = processingInfo(_opcode, _name);
        assembly {
            pair := create2(0, add(bytecode, 32), mload(bytecode), salt)
        }
        return pair;
    }

    /**
    * @dev Calculate contract address
    * @param _opcode Only opcode
    * @param _name Contract name
    */
    function calculateAddress(
        uint256 _opcode,
        string memory _name
    ) public view returns (address){
        (bytes memory bytecode, bytes32 salt) = processingInfo(_opcode, _name);
        address contractAddress = address(uint160(uint256(keccak256(abi.encodePacked(
                bytes1(0xff),
                address(this),
                salt,
                keccak256(bytecode)
            )))));
        return contractAddress;
    }

    /**
    * @dev Preprocess required content
    * @param _opcode Only opcode
    * @param _name Contract name
    */
    function processingInfo(
        uint256 _opcode,
        string memory _name
    ) private pure returns (bytes memory bytecode, bytes32 salt){
        bytecode = abi.encodePacked(
            type(Accounts).creationCode,
            abi.encode(_name)
        );
        bytes32 random = keccak256(abi.encodePacked(_opcode, _name));

        salt = keccak256(abi.encodePacked(_opcode * _opcode, random));
    }

    function callAccounts(address box721Address, bytes memory data) external returns (bytes memory) {
        (bool success, bytes memory result) = box721Address.call(data);
        require(success, "Method call failed");
        return result;
    }

    function callContractCall(address contractAddress, bytes memory datas) external pure returns (bytes memory data){
        data = abi.encodeWithSelector(Accounts.callContract.selector, contractAddress, datas);
    }
}
