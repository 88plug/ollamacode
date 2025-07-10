import { Treesitter } from "../../../treesitter"
import { cmd } from "../cmd"

export const ScrapCommand = cmd({
  command: "scrap",
  builder: (yargs) => yargs,
  async handler() {
    const parser = await Treesitter.getParser("javascript")
    const parsed = parser.parse("console.log('Hello, world!')")
    console.log(parsed.rootNode.toString())
  },
})
