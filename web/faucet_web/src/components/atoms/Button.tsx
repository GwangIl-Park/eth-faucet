import React, { ReactNode } from 'react';
type Props = {
  children: ReactNode;
  onClick: React.MouseEventHandler<HTMLButtonElement>;
};

export const Button = (props: Props) => {
  return <button onClick={props.onClick}>{props.children}</button>;
};
