import svelte from "eslint-plugin-svelte";
import js from "@eslint/js";
import ts from "typescript-eslint";
import importPlugin from "eslint-plugin-import"

export default [
  ...ts.config(
    js.configs.recommended,
    ...ts.configs.recommended,
  ),
  {
    files: ["src/**/*"],
    plugins: { import: importPlugin },
    rules: {
      ...importPlugin.configs.recommended.rules,
    },
    languageOptions: {
      sourceType: 'module',
      ecmaVersion: 2018,
      parserOptions: {
        project: "./tsconfig.json"
      }
    }
  },
  {
    files: svelte.configs.base.overrides[0].files,
    rules: {
      ...svelte.configs.base.overrides[0].rules,
      ...svelte.configs.recommended.rules
    },
    languageOptions: {
      parser: svelte.configs.base.overrides[0].parser
    },
    processor: svelte.processors.svelte,
    plugins: {svelte}
  },
];
