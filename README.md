
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


