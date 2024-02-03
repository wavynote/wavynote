type FolderType = "light" | "dark" | "disable" | "editable" | "focused";

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
    case "disable":
      return "disable";
    case "editable":
      return "editable";
    case "focused":
      return "focused";
  }
}

export default function FolderBtn({ name, type = "light", ...rest }: FolderProps) {
  
  return (
    <div className={`FolderBtn ${getTheme(type)}`} {...rest}>
      <button className="delBtn">삭제</button>
      <input type="text" value={name}></input>
      <p>노트 <span>0</span>개</p>
    </div>
  );
}