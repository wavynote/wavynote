type ButtonType = "light" | "dark" | "newNote";

interface ButtonProps {
  name: string;
  type: ButtonType;
  onClick?: () => void;
}

function getTheme(type: ButtonType) {
  switch (type) {
    case "light":
      return "default";
    case "dark":
      return "dark";
    case "newNote":
      return "dark newNote";
    default:
      return "";
  }
}

export default function Button({ name, type = "dark", onClick, ...rest }: ButtonProps) {
  const handleClick = () => {
    if( onClick ){
      onClick();
    }  
  }
  
  return (
    <button className={`textBtn ${getTheme(type)}`} onClick={handleClick} {...rest}>
      {name}
    </button>
  );
}
