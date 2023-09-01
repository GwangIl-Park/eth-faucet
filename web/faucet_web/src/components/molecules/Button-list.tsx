import { Button } from '../atoms/Button';
type Props = {
  onClickETH: React.MouseEventHandler<HTMLButtonElement>;
  onClickToken: React.MouseEventHandler<HTMLButtonElement>;
};

export const ButtonList = (props:Props) => {

  return (
    <div className="ButtonList">
      <Button onClick={props.onClickETH}>Send ETH</Button>
      <Button onClick={props.onClickToken}>Send Token</Button>
    </div>
  );
};
