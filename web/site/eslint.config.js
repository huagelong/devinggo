// eslint.config.js
import antfu from '@antfu/eslint-config'

export default antfu(
  {
    rules: {
      '@typescript-eslint/semi': 'off',
      'no-unused-vars': 'off',
      'no-undef': 'off',
      'no-debugger': 'warn',
      'no-console': 'warn',
      'curly': 'off',
      'node/prefer-global/process': 'off',
      'unused-imports/no-unused-vars': 'warn',
    },
  },
)
