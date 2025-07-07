import type { MessageV2 } from "opencode/session/message-v2"
import { Show, For, Switch, Match, createSignal, type JSX } from "solid-js"

import CodeBlock from "../CodeBlock"
import styles from "../share.module.css"

type ToolProps = {
  id: MessageV2.ToolPart["id"]
  tool: MessageV2.ToolPart["tool"]
  state: MessageV2.ToolStateCompleted
  rootDir?: string
  isLastPart?: boolean
}

interface TerminalPartProps extends JSX.HTMLAttributes<HTMLDivElement> {
  command: string
  error?: string
  result?: string
  desc?: string
  expand?: boolean
}


interface ResultsButtonProps extends JSX.HTMLAttributes<HTMLButtonElement> {
  showCopy?: string
  hideCopy?: string
  results: boolean
}
function ResultsButton(props: ResultsButtonProps) {
  return (
    <button type="button" data-element-button-text data-element-button-more {...props}>
      <span>{props.results ? props.hideCopy || "Hide results" : props.showCopy || "Show results"}</span>
    </button>
  )
}

interface TextPartProps extends JSX.HTMLAttributes<HTMLDivElement> {
  text: string
  expand?: boolean
}
function TextPart(props: TextPartProps) {
  return (
    <div class={styles["message-text"]} data-expanded={props.expand === true} {...props}>
      <pre>{props.text}</pre>
    </div>
  )
}

function TerminalPart(props: TerminalPartProps) {
  return (
    <div class={styles["message-terminal"]} data-expanded={props.expand === true} {...props}>
      <div data-section="body">
        <div data-section="header">
          <span>{props.desc}</span>
        </div>
        <div data-section="content">
          <CodeBlock lang="bash" code={props.command} />
          <Switch>
            <Match when={props.error}>
              <CodeBlock lang="text" data-section="error" code={props.error || ""} />
            </Match>
            <Match when={props.result}>
              <CodeBlock lang="console" code={props.result || ""} />
            </Match>
          </Switch>
        </div>
      </div>
    </div>
  )
}

export function BashTool(props: ToolProps) {
  const command = () => props.state.metadata?.title
  const desc = () => props.state.metadata?.description
  const result = () => props.state.metadata?.stdout
  const error = () => props.state.metadata?.stderr

  return (
    <>
      {command() && (
        <TerminalPart desc={desc()} data-size="sm" command={command()!} result={result()} error={error()} />
      )}
    </>
  )
}

export function FallbackTool(props: ToolProps) {
  const [showResults, setShowResults] = createSignal(false)

  return (
    <>
      <div data-part-title>{props.tool}</div>
      <div data-part-tool-args>
        <For each={flattenToolArgs(props.state.input)}>
          {(arg) => (
            <>
              <div></div>
              <div>{arg[0]}</div>
              <div>{arg[1]}</div>
            </>
          )}
        </For>
      </div>
      <Switch>
        <Match when={props.state.output}>
          <div data-part-tool-result>
            <ResultsButton results={showResults()} onClick={() => setShowResults((e) => !e)} />
            <Show when={showResults()}>
              <TextPart expand data-size="sm" data-color="dimmed" text={props.state.output} />
            </Show>
          </div>
        </Match>
      </Switch>
    </>
  )
}

// Converts nested objects/arrays into [path, value] pairs.
// E.g. {a:{b:{c:1}}, d:[{e:2}, 3]} => [["a.b.c",1], ["d[0].e",2], ["d[1]",3]]
function flattenToolArgs(obj: any, prefix: string = ""): Array<[string, any]> {
  const entries: Array<[string, any]> = []

  for (const [key, value] of Object.entries(obj)) {
    const path = prefix ? `${prefix}.${key}` : key

    if (value !== null && typeof value === "object") {
      if (Array.isArray(value)) {
        value.forEach((item, index) => {
          const arrayPath = `${path}[${index}]`
          if (item !== null && typeof item === "object") {
            entries.push(...flattenToolArgs(item, arrayPath))
          } else {
            entries.push([arrayPath, item])
          }
        })
      } else {
        entries.push(...flattenToolArgs(value, path))
      }
    } else {
      entries.push([path, value])
    }
  }

  return entries
}
