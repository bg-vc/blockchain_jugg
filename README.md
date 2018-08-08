
NODE 3000:（中心节点，所有其他节点都会连接到这个节点，这个节点会在其他节点之间发送数据）

export NODE_ID=3000

./blockchain createwallet

./blockchain listaddresses

1DzSseaG5P6XfqACDT7dWUgeM9p9ghG2uA

18wHKBZYJrkDHPw9RJmwubcFHhGNZ7iuxv

./blockchain createblockchain -address 1DzSseaG5P6XfqACDT7dWUgeM9p9ghG2uA

./blockchain getbalance -address 1DzSseaG5P6XfqACDT7dWUgeM9p9ghG2uA

cp blockchain_3000.db blockchain_genesis.db 


NODE 3001:（钱包节点，转账交易） 

export NODE_ID=3001

./blockchain createwallet

./blockchain listaddresses

1D1UPZa6mLbiUR1coqt6U4zRBxL4fSjkBC

16zkeVpwXYgbigKvxabaVtu5yUMv9f4Jb

cp blockchain_genesis.db blockchain_3001.db


NODE 3000:

./blockchain send -from 1DzSseaG5P6XfqACDT7dWUgeM9p9ghG2uA -to 1D1UPZa6mLbiUR1coqt6U4zRBxL4fSjkBC -amount 10 -mine

./blockchain send -from 1DzSseaG5P6XfqACDT7dWUgeM9p9ghG2uA -to 16zkeVpwXYgbigKvxabaVtu5yUMv9f4Jb -amount 10 -mine

./blockchain getbalance -address 1DzSseaG5P6XfqACDT7dWUgeM9p9ghG2uA

./blockchain startnode

NODE 3001:

./blockchain startnode

./blockchain getbalance -address 1D1UPZa6mLbiUR1coqt6U4zRBxL4fSjkBC

./blockchain getbalance -address 16zkeVpwXYgbigKvxabaVtu5yUMv9f4Jb

NODE 3002:(挖矿节点，在内存池中存储新的交易，当有足够的交易时，打包挖出一个新块)

export NODE_ID=3002

./blockchain createwallet

1479vPBbMG6k8PjJ75Bt1SikzVKefLj5A8

./blockchain startnode -miner 1479vPBbMG6k8PjJ75Bt1SikzVKefLj5A8

NODE 3001:
./blockchain send -from 1D1UPZa6mLbiUR1coqt6U4zRBxL4fSjkBC -to 1479vPBbMG6k8PjJ75Bt1SikzVKefLj5A8 -amount 2

./blockchain send -from 16zkeVpwXYgbigKvxabaVtu5yUMv9f4Jb -to 1479vPBbMG6k8PjJ75Bt1SikzVKefLj5A8 -amount 3

./blockchain startnode

./blockchain getbalance -address 1D1UPZa6mLbiUR1coqt6U4zRBxL4fSjkBC

./blockchain getbalance -address 16zkeVpwXYgbigKvxabaVtu5yUMv9f4Jb

./blockchain printchain


NODE 3002:
./blockchain getbalance -address 1479vPBbMG6k8PjJ75Bt1SikzVKefLj5A8


=========== Block 000064b3c9d7e5431546e8bfdc3e36affd786d5be2ddbf3bc3936f51c566f985 ============

Height: 5

Prev. block: 00005bfaf8da7aeb02cbcb9fb58fff31d82ab046143f32f94862c5dc4ce5aec2

PoW: true

Transaction 5f07e908e208e5c206f9d5864f329f8c586fd57e0b37266df8bd0364b4a115da:

  Input 0:
  
      TXID:  
      Out:       -1
      Signature: 
      PubKey:    66663937386166383863373761643132623239366166663538373964366235653333393463323139
  
  Output 0:
  
      Index:  0
      Value:  50
      Script: 42eb829367dbfdce3ce2c752e3afa2b6a2e494e0
Transaction 5475afe5af5a9dcd0f17dfb22ce43f16b049ee15d22e5cc0f25d65da2da65ea8:

  Input 0:
  
      TXID:      41332e2898b0990a44180f379849dad192d2d27c7d4f9719d19b98cfac55a928
      Out:       0
      Signature: 44e3aca1e16eb6d744c5815d9a700b2b256a38c47f59ece9a58b49740154ab206ea5cc48e5069debe8b6422f8c656ab35823d6b046f0b51f93b3ae5f8e11f348
      PubKey:    f24c48983c2e883cdd7eedca71712d230b8af2f5dbb2fc637e8ec95cedefb32febe21cc3139c2c35f5fc1c278f20cd705cef7d858c7f2a8965c9d94cda2a1835
  
  Input 1:
  
      TXID:      903a07c1b0996d6ff4a43b2cf583b81db656357b9220f9a7770aa7aece7ff0d5
      Out:       1
      Signature: b9441f04e4f8802097c899eddae664a20c344361d8dcad55034f96c420467be9f3b9736f374b68d76dcd6a60f03e10faf1882cdb24525fd44707756c973e7082
      PubKey:    f24c48983c2e883cdd7eedca71712d230b8af2f5dbb2fc637e8ec95cedefb32febe21cc3139c2c35f5fc1c278f20cd705cef7d858c7f2a8965c9d94cda2a1835
 
 Input 2:
 
      TXID:      a26cf9f3d4628115d2f1318e277abb114ef9122901ef3b189382478da47c5293
      Out:       0
      Signature: 973cd5240e67ea7bedd9fb1ea700f20c28a58cdced2711e5e1dd66985c23010d940bf4b652fcb9d73e4d23423315067dc424ac9d0c2c45a8d421e391a958215d
      PubKey:    f24c48983c2e883cdd7eedca71712d230b8af2f5dbb2fc637e8ec95cedefb32febe21cc3139c2c35f5fc1c278f20cd705cef7d858c7f2a8965c9d94cda2a1835
  
  Input 3:
  
      TXID:      d4a7a7ac40ab91d3c1462e850c6bcba109bfaffe94028af8f30f2ab1dfc53faa
      Out:       0
      Signature: b478a42e5a5a928f933272644da7733be5323e888d02da39e652c8bcbf6570f23eab261544a3f80b5c41a2a7b7e04b750ce91a4495ecfe435e954a7cb9805bc8
      PubKey:    f24c48983c2e883cdd7eedca71712d230b8af2f5dbb2fc637e8ec95cedefb32febe21cc3139c2c35f5fc1c278f20cd705cef7d858c7f2a8965c9d94cda2a1835
  
  Output 0:
  
      Index:  0
      Value:  165
      Script: f6a6a4fc30fd0147eab078b510b4bb3d176a90f1
 
 Output 1:
 
      Index:  1
      Value:  5
      Script: 42eb829367dbfdce3ce2c752e3afa2b6a2e494e0


============ Block 00005bfaf8da7aeb02cbcb9fb58fff31d82ab046143f32f94862c5dc4ce5aec2 ============

Height: 4

Prev. block: 000007340d52d180e86889cfa602e6cbfa3f27fa370e44e542902d19191e4b49

PoW: true

Transaction 41332e2898b0990a44180f379849dad192d2d27c7d4f9719d19b98cfac55a928:

  Input 0:
  
      TXID:      
      Out:       -1
      Signature: 
      PubKey:    35393536373063656662366438346363663163633963353135653433353435656266336532393563
 
 Output 0:
      Index:  0
      Value:  50
      Script: 42eb829367dbfdce3ce2c752e3afa2b6a2e494e0
      
Transaction 903a07c1b0996d6ff4a43b2cf583b81db656357b9220f9a7770aa7aece7ff0d5:

  Input 0:
  
      TXID:      6c308b68e8780cc093f32342ceef3d841b4abecbd4d0e8e0edd477b285e8c67e
      Out:       1
      Signature: 4c114f89f56b50740fcb615cc9580ab704da867b4af478d707885618d9e3f677dce9b941dce9a4f07622b646fca804d308cd0bb0cc02af8867275da1f7a37009
      PubKey:    f24c48983c2e883cdd7eedca71712d230b8af2f5dbb2fc637e8ec95cedefb32febe21cc3139c2c35f5fc1c278f20cd705cef7d858c7f2a8965c9d94cda2a1835
  
  Output 0:
  
      Index:  0
      Value:  10
      Script: f6a6a4fc30fd0147eab078b510b4bb3d176a90f1
  
  Output 1:
  
      Index:  1
      Value:  20
      Script: 42eb829367dbfdce3ce2c752e3afa2b6a2e494e0


============ Block 000007340d52d180e86889cfa602e6cbfa3f27fa370e44e542902d19191e4b49 ============

Height: 3

Prev. block: 00004c0025c5e2517e5b796ce8dfe3a980eff888aa24d2bc226f3fcc81a7958f

PoW: true

Transaction d4a7a7ac40ab91d3c1462e850c6bcba109bfaffe94028af8f30f2ab1dfc53faa:
  Input 0:
      TXID:      
      Out:       -1
      Signature: 
      PubKey:    34636135306465363232343833326233336430333839656432616566363265353266613635666537
  Output 0:
      Index:  0
      Value:  50
      Script: 42eb829367dbfdce3ce2c752e3afa2b6a2e494e0
      
Transaction 6c308b68e8780cc093f32342ceef3d841b4abecbd4d0e8e0edd477b285e8c67e:

  Input 0:
  
      TXID:      1290556b2401320b962b7938ab72f0d4b98e0b23c016769d8044c44ed5509361
      Out:       1
      Signature: 24152ab676520e8cec7320c411cca7305fd06ebf65e8fbf908f33c4871d0074ef5679e0b0e59781005bf1ea9c03a1c01afe3b860f322e357a32ae20520cc60b9
      PubKey:    f24c48983c2e883cdd7eedca71712d230b8af2f5dbb2fc637e8ec95cedefb32febe21cc3139c2c35f5fc1c278f20cd705cef7d858c7f2a8965c9d94cda2a1835
 
 Output 0:
 
      Index:  0
      Value:  10
      Script: 05ac83e236430288c4cc1e4bd016d183cbb372a1
  
  Output 1:
  
      Index:  1
      Value:  30
      Script: 42eb829367dbfdce3ce2c752e3afa2b6a2e494e0


============ Block 00004c0025c5e2517e5b796ce8dfe3a980eff888aa24d2bc226f3fcc81a7958f ============

Height: 2

Prev. block: 0000c9401cd89b06a81625a3fef631c355738485be4c7ded5051b9dfad9dd447

PoW: true

Transaction a26cf9f3d4628115d2f1318e277abb114ef9122901ef3b189382478da47c5293:

  Input 0:
  
      TXID:      
      Out:       -1
      Signature: 
      PubKey:    35313264666261326561373739633438626161613462316561313663326638363138653238323739
 
 Output 0:
 
      Index:  0
      Value:  50
      Script: 42eb829367dbfdce3ce2c752e3afa2b6a2e494e0
      
Transaction 1290556b2401320b962b7938ab72f0d4b98e0b23c016769d8044c44ed5509361:

  Input 0:
  
      TXID:      5353ac882690dd7366d234423402412d5af3a3708ca2f06f3f36a89eaab14726
      Out:       1
      Signature: a35ae7b044c807a0a432816142c1b57687b423c56d989f0c832d34e517ae7df04b80957df3fb715c999f0ad6ce61f423e88547e4c481f681ce27a634ef951f2b
      PubKey:    f24c48983c2e883cdd7eedca71712d230b8af2f5dbb2fc637e8ec95cedefb32febe21cc3139c2c35f5fc1c278f20cd705cef7d858c7f2a8965c9d94cda2a1835
 
 Input 1:
 
      TXID:      836c42cc2cf586e4b0d2cd3e6155f6d24ec0766b54ace2d6def3b2e66c8523b2
      Out:       0
      Signature: 8fc55f0e50bd5d37e84276d137336d193c6a4be1f5e9a973989b6d97a1dbb36eaeddb32d614f882664a4d10d822c1ad9a36cbe599a424ab75f5c932265e13206
      PubKey:    f24c48983c2e883cdd7eedca71712d230b8af2f5dbb2fc637e8ec95cedefb32febe21cc3139c2c35f5fc1c278f20cd705cef7d858c7f2a8965c9d94cda2a1835
 
 Output 0:
 
      Index:  0
      Value:  30
      Script: 05ac83e236430288c4cc1e4bd016d183cbb372a1
 
 Output 1:
 
      Index:  1
      Value:  40
      Script: 42eb829367dbfdce3ce2c752e3afa2b6a2e494e0


============ Block 0000c9401cd89b06a81625a3fef631c355738485be4c7ded5051b9dfad9dd447 ============

Height: 1

Prev. block: 0000c53daae0b4b0cdb3bec8890ecdadf40e6dc5e66d9162dc5161110ed779d7

PoW: true

Transaction 836c42cc2cf586e4b0d2cd3e6155f6d24ec0766b54ace2d6def3b2e66c8523b2:

  Input 0:
  
      TXID:      
      Out:       -1
      Signature: 
      PubKey:    66316265386232366361376639396532623864386336356630313063633066386533353839633861
  
  Output 0:
  
      Index:  0
      Value:  50
      Script: 42eb829367dbfdce3ce2c752e3afa2b6a2e494e0
      
Transaction 5353ac882690dd7366d234423402412d5af3a3708ca2f06f3f36a89eaab14726:

  Input 0:
  
      TXID:      8a26cde23205176e92938fd0133d726ba521307203e6693a305277e40ea5f2c7
      Out:       0
      Signature: 937349db1c99cb190758b25007671ac2eb16f383e2ba26d5a1ae160e0485ee02308b9a760e572694fa159d2d6f1675cc50f38d5b2bf01d19cb0bc045e0cf349e
      PubKey:    f24c48983c2e883cdd7eedca71712d230b8af2f5dbb2fc637e8ec95cedefb32febe21cc3139c2c35f5fc1c278f20cd705cef7d858c7f2a8965c9d94cda2a1835
  
  Output 0:
  
      Index:  0
      Value:  30
      Script: f6a6a4fc30fd0147eab078b510b4bb3d176a90f1
  
  Output 1:
  
      Index:  1
      Value:  20
      Script: 42eb829367dbfdce3ce2c752e3afa2b6a2e494e0


============ Block 0000c53daae0b4b0cdb3bec8890ecdadf40e6dc5e66d9162dc5161110ed779d7 ============

Height: 0

Prev. block: 

PoW: true

Transaction 8a26cde23205176e92938fd0133d726ba521307203e6693a305277e40ea5f2c7:

  Input 0:
  
      TXID:      
      Out:       -1
      Signature: 
      PubKey:    5468652054696d65732030332f4a616e2f32303039204368616e63656c6c6f72206f6e206272696e6b206f66207365636f6e64206261696c6f757420666f722062616e6b73
  
  Output 0:
  
      Index:  0
      Value:  50
      Script: 42eb829367dbfdce3ce2c752e3afa2b6a2e494e0


