type FolderType = "light" | "dark" ;

interface FolderProps {
  name: string;
  type: FolderType;
}

function getTheme(type: FolderType) {
  switch (type) {
    case "light":
      return "light";
    case "dark":
      return "dark";
  }
}

export default function FolderBtn({ name, type = "light", ...rest }: FolderProps) {
  return (
    <div className={`FolderBtn ${getTheme(type)}`} {...rest}>
      {name}
      <p>노트 <span>0</span>개</p>
    </div>
  );
}