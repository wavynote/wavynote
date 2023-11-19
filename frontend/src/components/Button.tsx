type ButtonTheme = 'light' | 'dark';

interface ButtonProps {
  name: string;
  theme: ButtonTheme;
}

function getTheme(theme: ButtonTheme) {
  switch (theme) {
    case 'light':
      return 'bg-white text-zinc-800';
    case 'dark':
      return 'bg-gray-950 text-white';
  }
}

export default function Button({ name, theme = 'dark', ...rest }: ButtonProps) {
  return (
    <button
      className={`px-12 py-4 rounded-[20px] text-sm font-light ${getTheme(
        theme,
      )} shadow-[3px_0px_20px_5px_rgba(56,79,138,0.05)]`}
      {...rest}
    >
      {name}
    </button>
  );
}
