type Props = {
  name: string;
};

export const Label = (props: Props) => {
  return (
      <label>{props.name}</label>
  );
};