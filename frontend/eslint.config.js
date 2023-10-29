import eslintConfigPrettier from 'eslint-config-prettier';

module.exports = [
  {
    files: ['**/*.ts', '**/*.tsx'],
    ignores: ['**/*.config.js', '**/*.config.ts'],
    rules: {
      semi: 'error',
    },
  },
  eslintConfigPrettier,
];
