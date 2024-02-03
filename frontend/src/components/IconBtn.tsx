type ButtonType = "prev" | "search" | "searchBlack" | "etc";

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
    case "searchBlack":
      return "searchBlackBtn";
    case "etc":
      return "etcBtn"
  }
}

export default function Button({ name, type = "prev", ...rest }: ButtonProps) {
  return (
    <button className={`icBtn ${getTheme(type)}`} {...rest}>
      {name}
    </button>
  );
}
