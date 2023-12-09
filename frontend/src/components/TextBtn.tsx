type ButtonType = "light" | "dark" | "newNote";

interface ButtonProps {
  name: string;
  type: ButtonType;
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
