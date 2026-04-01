# Frontend Architecture Rules — PaaS Vercel Clone

## 🏗 Layers (in exact order of dependency)
`app/` → `controllers/` → `usecases/` → `interfaces/` ← `services/` ← `di/`

---

## 🚦 Controller Rules
- **Always a TypeScript Class**.
- Receives services via DI Container (invoked via `getContainer()`). Do not inject via file imports.
- Orchestrates Use Cases and handles data validation using `Zod` before passing input down.

## 🧠 Use Case Rules
- One Use Case class = **One Job** (Single Responsibility Principle).
- Constructor must accept injected interfaces, never concrete services.
- Return domain models (from `types/`), not raw HTTP responses.

## 🧱 Component Rules
- Use Server Components where possible, except when consuming Hooks.
- Custom logic hooks (`hooks/`) wrap the Controllers.
- **Scoping**:
  - Used in one page only: `app/[page]/_components/`
  - Shared across multiple pages: `components/shared/`

## 📡 Service Rules
- Only layer allowed to perform `fetch` or `axios` calls (`lib/api.ts`).
- Concrete service must `implements` an interface from `interfaces/`.

---

## 📝 Update Flow Reminder
After adding new code, modules, or completing a feature:
1. Update the `context.md` in the modified folder.
2. Read this `prompt.md` to ensure you adhered to the architecture rules.
