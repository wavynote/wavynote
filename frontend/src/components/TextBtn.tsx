import Link from "next/link";

type ButtonType = "light" | "dark" | "newNote";

interface ButtonProps {
  name: string;
  type: ButtonType;  
  component: React.ElementType;
}

function getTheme(type: ButtonType) {
  switch (type) {
    case "light":
      return "default";
    case "dark":
      return "dark";
    case "newNote":
      return "dark newNote"
  }
}

export default function Button({ name, type = "dark", ...rest }: ButtonProps) {
  return (
    <button className={`textBtn ${getTheme(type)}`} {...rest}>
      {name}
    </button>
  );
}
