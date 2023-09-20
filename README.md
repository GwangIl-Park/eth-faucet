
# 전체 구조
 
![image](https://github.com/GwangIl-Park/eth-faucet/assets/40749130/d0ddc736-c438-4edb-84d4-299bd6a4c629)

## API

|  | URL | method |
| --- | --- | --- |
| ETH 요청 | /faucetETH/request | POST |
| Token 요청 | /faucetToken/request | POST |

## Data

| type | Data |
| --- | --- |
| Request | {wallet_address: “사용자 address”} |
| Response | {transaction_hash: “transaction Hash값”eth_balance: “사용자의 현재 eth양”token_balance: “사용자의 현재 token양”}|

## DB

| name | type | description |
| --- | --- | --- |
| address | string | account address |
| amount | string | 총 전송 받은 양 |
| time | time.Time | 마지막 요청 성공 시간 |

# 실행

## 서버

- 빌드 후 실행

```jsx
go build
./eth-faucet
```

- flags

![image](https://github.com/GwangIl-Park/eth-faucet/assets/40749130/f8da5155-b008-48ec-9434-0135087b0c5a)


## Front

- in web/faucet-web

```jsx
npm run start
```

- 실행 화면 확인

![image](https://github.com/GwangIl-Park/eth-faucet/assets/40749130/bcc907f3-afd3-4db1-958c-e8ad69f769f4)

# 3. 기능 확인

---

## ETH 전송

### 성공 케이스

- 최초 전송
    - 서버 (log level : debug)

    ![image](https://github.com/GwangIl-Park/eth-faucet/assets/40749130/b3eb4473-f280-49f9-87a5-c46c44cfd779)



    
    - client
    

    
    - 노드 balance 확인
        
        ![image](https://github.com/GwangIl-Park/eth-faucet/assets/40749130/0b90967b-42ca-4531-8f67-eb831b6fd256)

        

- delay 기간 지난 이후 요청
    - 서버 (log level : debug)

![image](https://github.com/GwangIl-Park/eth-faucet/assets/40749130/0ed64a26-3473-4bbb-acda-53fe8eb8eb45)


### 실패 케이스

- delay 기간 이내에 요청
    - 서버 (log level : debug)
    
    ![image](https://github.com/GwangIl-Park/eth-faucet/assets/40749130/d7209d02-2908-4b05-ae7e-aec7871c6ca1)

    
    - client (Error message에 담아서 전송)
    
    ![image](https://github.com/GwangIl-Park/eth-faucet/assets/40749130/ea8a62e7-4a29-4233-ad98-935760ddff3c)
