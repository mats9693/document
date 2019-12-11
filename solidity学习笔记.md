# solidity学习笔记
> solidity语言学习和使用的个人经验分享。

---
### 学习资料
http://www.tryblockchain.org/ 基础知识  
http://remix.ethereum.org 在线IDE，支持几乎全部solidity版本的编译、部署和调试  
https://solidity-cn.readthedocs.io/zh/develop/miscellaneous.html#index-7 中文文档速查表  
https://blog.csdn.net/wcc19840827/article/details/87534778 智能合约优化方向  

--- 
### solidity入门
solidity上手很简单，看一遍官方文档的例子，或者只是看一遍基础知识就可以开始写代码了  
solidity很多地方如果光靠分析就想得到准确的结论，那需要掌握大量的EVM、solidity底层实现等知识；所以另外的一种方法就是试，在remix上使用玩具代码测试思路，不仅可以根据结果指导实践，一些时候还能从结果中展开推论（remix支持transaction调试）

---
### trivia
1. 每个字符（char）占1个字节（Byte）的存储空间  
1. 定长数组最长为32个字符，即bytes32  
1. solidity不支持不定长数组的数组的解析和传递（函数输入输出参数位置不能出现string[]）  
1. 不定长数组可以使用.length属性和.push函数进行操作：
    ```solidity  
    uint256[] ia;
    
    ia.push(10);
    ia.push(20);
    ia.push(30);
            
    (ia[0], ia[ia.length-1]) = (ia[ia.length-1], ia[0]);
            
    ia.length--;
    ia.push(40);
    ```  
1. struct默认使用internal修饰，不能修改；即，只有定义了struct的合约、继承了定义合约的合约能够访问该struct，单纯的import不行（lib除外）
1. solidity使用代码复制的方式实现继承：
    ```solidity  
    contract A {
        int a_int;
    }
    
    contract isA is A {
        function getAInt() public view returns(int) {
            return a_int;
        }
    }
    
    
    //  A + isA above is the same as A bellow. 
    
    
    contract A {
        int a_int;
        
        function getAInt() public view returns(int) {
            return a_int;
        }
    }
    ```
1. 合约之间的调用：msg.sender / context
1. A合约调用B合约函数的写法：  
   （底层使用.call实现，改变msg.sender和context）  
   （使用.delegatecall方法，不改变msg.sender，改变context  (lib)）  
   （继承和lib除外，它们不属于合约之间的调用）  
    ```solidity  
    contract A {
      B bIns;
      constructor (address bAddr) {
    bIns = B(bAddr);
    }
    bIns.funcInB ... ...;
    }
    
    // what is context main about? 
    contract A {
      int  i;
    
      function setInt(int _value) public {
    i = _value;
    }
    }
    contract B {
      int  i;
    
      function setInt(int _value) public {
    i = _value;
    }
    }
    
    // library:
    library L {
      function lFunc() {}    // ‘L.lFunc();’ in A， like call cross contract, but not. 
    }
    
    // delegatecall, notice ‘int8’
    contract A {
        int i = 1;
        
        function testFunc(address addr) public returns (bool) {
            return addr.delegatecall(bytes4(keccak256("testFuncB(int8)")), 10);
        }
    }
    
    contract B {
        int i = 2;
        
        event ShowParams(address msgSender, address contractAddr, int _i);
        
        function testFuncB(int8 _value) public {
            if (_value == 10) {
                emit ShowParams(msg.sender, address(this), i);    
            }
        }
    }
    ```
1. 调用其他合约的方法时，如果目标合约没有对应方法，则调用fallback函数（一个匿名函数）。
1. truffle：一个solidity开发框架  
   - truffle可以进行solidity编译、智能合约部署、智能合约测试（js、solidity）、transaction调试、区块链数据查询（balance、blockNumber）等操作  
   - truffle需要与一条链一起工作，可以使用ganache / geth搭建一条链，使用geth搭建一条私链，在dp项目中有相关脚本，windows和linux的都有
1. event：事件
    ```solidity  
    event RegisterVerifier(string seqNo, address[] users);
    ```
   - 智能合约的事件就像是其他程序中的写日志。
   - 智能合约会在emit的位置，为事件的参数赋值，然后把事件写到一个交易的log里面。

---
###### Mario
I've been pretending to work hard, but you're really growing up.
