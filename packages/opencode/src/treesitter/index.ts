import Parser from "tree-sitter"
import Javascript from "tree-sitter-javascript"
import { App } from "../app/app"

export namespace Treesitter {
  const parsers = new Map<string, Parser>()

  const LANGUAGES: Record<string, Parser.Language> = {
    javascript: Javascript.language as any,
  }

  export async function getParser(language: string) {
    const match = parsers.get(language)
    if (match) return match
    const parser = new Parser()
    parser.setLanguage(LANGUAGES[language])
    parsers.set(language, parser)
    return parser
  }

  export async function init() {
    const app = App.info()
    if (!app.git) return
  }
}
