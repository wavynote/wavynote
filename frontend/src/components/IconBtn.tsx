type ButtonType = "prev" | "search";

interface ButtonProps {
  name: string;
  type: ButtonType;
}

function getTheme(type: ButtonType) {
  switch (type) {
    case "prev":
      return "prevBtn";
    case "search":
      return "searchBtn";
  }
}

export default function Button({ name, type = "prev", ...rest }: ButtonProps) {
  return (
    <button className={`icBtn ${getTheme(type)}`} {...rest}>
      {name}
    </button>
  );
}
