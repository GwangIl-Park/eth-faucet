import { useEffect, useState } from 'react';
import { Input } from '../atoms/Input';
import { ButtonList } from '../molecules/Button-list';
import axios from 'axios';

export const FaucetContent = () => {
  const instance = axios.create({
    baseURL: process.env.REACT_APP_FAUCET_SERVER_URL,
  });

  instance.defaults.withCredentials = true;

  const [address, setAddress] = useState('');
  const [ethAmount, setEthAmount] = useState('');
  const [tokenAmount, setTokenAmount] = useState('');
  const [showText, setShowText] = useState(false);
  const [isLoading, setIsLoading] = useState(false);

  useEffect(() => {
    const url = new URL(window.location.href);
    const urlParams = url.searchParams;
    const addressParam = urlParams.get('address');

    setAddress(addressParam ? addressParam : '');
  }, []);

  const onChange = (event: any) => {
    setAddress(event.target.value);
  };

  const onClickBtn = () => {
    setIsLoading(true);
    instance
      .post('/faucet/request', {
        wallet_address: address,
      })
      .then((response) => {
        setEthAmount((Number(response.data.ethBalance) / 10 ** 18).toFixed(6));
        setTokenAmount(
          (Number(response.data.tokenBalance) / 10 ** 18).toFixed(6)
        );
        setShowText(true);
        setIsLoading(false);
        alert(
          `Send Success\neth txHash : ${response.data.ethTransactionHash}\nToken txHash : ${response.data.tokenTransactionHash}`
        );
      })
      .catch((error) => {
        setIsLoading(false);
        alert(`Send Fail\nMessage : ${error.response.data.message}`);
      });
  };

  const onClickETH = () => {
    setIsLoading(true);
    instance
      .post('/faucetETH/request', {
        wallet_address: address,
      })
      .then((response) => {
        setEthAmount(response.data.ethBalance);
        setTokenAmount(response.data.tokenBalance);
        setShowText(true);
        setIsLoading(false);
        alert(`Send ETH Success\ntxHash : ${response.data.transactionHash}`);
      })
      .catch((error) => {
        setIsLoading(false);
        alert(`Send ETH Fail\nMessage : ${error.response.data.message}`);
      });
  };

  const onClickToken = () => {
    setIsLoading(true);
    instance
      .post('/faucetToken/request', {
        wallet_address: address,
      })
      .then((response) => {
        setEthAmount(response.data.ethBalance);
        setTokenAmount(response.data.tokenBalance);
        setShowText(true);
        setIsLoading(false);
        alert(`Send Token Success\ntxHash : ${response.data.transactionHash}`);
      })
      .catch((error) => {
        setIsLoading(false);
        alert(`Send Token Fail\nMessage : ${error.response.data.message}`);
      });
  };

  return (
    <div className="faucet-container">
      <div className="welcome">Gipark Faucet</div>
      <Input
        name="Address"
        type="string"
        onChange={onChange}
        placeholder="Type ETH Wallet Address Here"
        value={address}
      />
      {isLoading ? (
        <p>신청 중입니다..</p>
      ) : (
        <ButtonList onClickETH={onClickETH} onClickToken={onClickToken} />
      )}
      <br />
      {showText ? (
        <div className="balance-container">
          <span>ETH : </span>
          <span>{ethAmount}</span>
          <span> </span>
          <span>Token : </span>
          <span>{tokenAmount}</span>
        </div>
      ) : null}
    </div>
  );
};
