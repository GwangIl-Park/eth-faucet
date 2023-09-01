import React from 'react';
type Props = {
  name: string;
  type: string;
  placeholder: string
  onChange: React.ChangeEventHandler<HTMLInputElement>;
  value: string | null;
};

export const Input = (props: Props) => {
  return (
      <input name={props.name} type={props.type} onChange={props.onChange} placeholder={props.placeholder} value={props.value ? props.value : ""} />
  );
};