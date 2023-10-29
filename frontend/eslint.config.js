import eslintConfigPrettier from 'eslint-config-prettier';

export default [
  {
    files: ['**/*.ts', '**/*.tsx'],
    ignores: ['**/*.config.js', '**/*.config.ts'],
    rules: {
      semi: 'error',
    },
  },
  eslintConfigPrettier,
];
