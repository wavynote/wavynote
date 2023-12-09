/** @type {import('eslint').Linter.FlatConfig[]} */
import nextPlugin from "@next/eslint-plugin-next";
import reactPlugin from "eslint-plugin-react";
import hooksPlugin from "eslint-plugin-react-hooks";
import tsPlugin from "@typescript-eslint/eslint-plugin";
import tsParser from "@typescript-eslint/parser";
import eslintConfigPrettier from "eslint-config-prettier";
import importPlugin from "eslint-plugin-import";

export default [
  {
    plugins: {
      react: reactPlugin,
    },
    rules: {
      ...reactPlugin.configs["jsx-runtime"].rules,
    },
    settings: {
      react: {
        version: "detect",
      },
    },
  },
  {
    plugins: {
      "react-hooks": hooksPlugin,
    },
    rules: hooksPlugin.configs.recommended.rules,
  },
  {
    plugins: {
      "@next/next": nextPlugin,
    },
    rules: {
      ...nextPlugin.configs.recommended.rules,
      ...nextPlugin.configs["core-web-vitals"].rules,
    },
  },
  {
    ignores: [".next/*"],
  },
  {
    files: ["**/*.ts", "**/*.tsx"],
    ignore: ["**/*.config.js"],
    languageOptions: {
      parser: tsParser,
      parserOptions: {
        ecmaFeatures: { modules: true },
        ecmaVersion: "latest",
        project: "./tsconfig.json",
      },
    },
    plugins: {
      import: importPlugin,
      "@typescript-eslint": tsPlugin,
      tsPlugin,
    },
    settings: {
      "import/parsers": {
        espree: [".js", ".cjs", ".mjs", ".jsx"],
        "@typescript-eslint/parser": [".ts", ".tsx"],
      },
      "import/resolver": {
        typescript: true,
        node: true,
      },
    },
    rules: {
      ...tsPlugin.configs["eslint-recommended"].rules,
      ...tsPlugin.configs["recommended"].rules,
      // ...importPlugin.configs['recommended'].rules,
      // 'ts/return-await': 2,
      "import/order": [
        "error",
        {
          groups: ["builtin", "external", "internal"],
          pathGroups: [
            {
              pattern: "react",
              group: "builtin",
              position: "after",
            },
            {
              pattern: "react-dom",
              group: "builtin",
              position: "after",
            },
            {
              pattern: "*next/**",
              group: "builtin",
              position: "after",
              patternOptions: {
                partial: true,
              },
            },
          ],
          pathGroupsExcludedImportTypes: [],
          "newlines-between": "always",
          alphabetize: {
            order: "asc",
            caseInsensitive: true,
          },
        },
      ],
    },
  },
  eslintConfigPrettier,
];
