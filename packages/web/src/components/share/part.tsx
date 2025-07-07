import { createMemo, createSignal, For, Match, Show, Switch, type JSX, type ParentProps } from "solid-js"
import { IconCheckCircle, IconChevronDown, IconChevronRight, IconHashtag, IconSparkles } from "../icons"
import styles from "./part.module.css"
import type { MessageV2 } from "opencode/session/message-v2"
import { ContentText } from "./content-text"
import { ContentMarkdown } from "./content-markdown"
import { DateTime } from "luxon"
import CodeBlock from "../CodeBlock"
import map from "lang-map"
import type { Diagnostic } from "vscode-languageserver-types"
import { FallbackTool } from "./tool"
import { ContentCode } from "./content-code"
import { ContentDiff } from "./content-diff"

export interface PartProps {
  index: number
  message: MessageV2.Info
  part: MessageV2.AssistantPart | MessageV2.UserPart
  last: boolean
}

export function Part(props: PartProps) {
  const [copied, setCopied] = createSignal(false)
  const id = createMemo(() => props.message.id + "-" + props.index)

  return (
    <div
      class={styles.root}
      id={id()}
      data-component="part"
      data-type={props.part.type}
      data-role={props.message.role}
      data-copied={copied() ? true : undefined}
    >
      <div data-component="decoration">
        <div data-slot="anchor" title="Link to this message">
          <a
            href={`#${id()}`}
            onClick={(e) => {
              e.preventDefault()
              const anchor = e.currentTarget
              const hash = anchor.getAttribute("href") || ""
              const { origin, pathname, search } = window.location
              navigator.clipboard
                .writeText(`${origin}${pathname}${search}${hash}`)
                .catch((err) => console.error("Copy failed", err))

              setCopied(true)
              setTimeout(() => setCopied(false), 3000)
            }}
          >
            <Switch>
              <Match when={true}>
                <IconSparkles width={18} height={18} />
              </Match>
            </Switch>
            <IconHashtag width={18} height={18} />
            <IconCheckCircle width={18} height={18} />
          </a>
          <span data-slot="tooltip">Copied!</span>
        </div>
        <div data-slot="bar"></div>
      </div>
      <div data-component="content">
        {props.message.role === "user" && props.part.type === "text" && (
          <>
            <ContentText text={props.part.text} expand={props.last} /> <Spacer />
          </>
        )}
        {props.message.role === "assistant" && props.part.type === "text" && (
          <>
            <ContentMarkdown expand={props.last} text={props.part.text} />
            {props.last && props.message.role === "assistant" && props.message.time.completed && (
              <Footer
                title={DateTime.fromMillis(props.message.time.completed).toLocaleString(
                  DateTime.DATETIME_FULL_WITH_SECONDS,
                )}
              >
                {DateTime.fromMillis(props.message.time.completed).toLocaleString(DateTime.DATETIME_MED)}
              </Footer>
            )}
            <Spacer />
          </>
        )}
        {props.part.type === "step-start" && props.message.role === "assistant" && (
          <div data-component="step-start">
            <div data-slot="provider">{props.message.providerID}</div>
            <div data-slot="model">{props.message.modelID}</div>
          </div>
        )}
        {props.part.type === "tool" &&
          props.part.state.status === "completed" &&
          props.message.role === "assistant" && (
            <div data-component="tool" data-tool={props.part.tool}>
              <Switch>
                <Match when={props.part.tool === "grep"}>
                  <GrepTool
                    message={props.message}
                    id={props.part.id}
                    tool={props.part.tool}
                    state={props.part.state}
                  />
                </Match>
                <Match when={props.part.tool === "glob"}>
                  <GlobTool
                    message={props.message}
                    id={props.part.id}
                    tool={props.part.tool}
                    state={props.part.state}
                  />
                </Match>
                <Match when={props.part.tool === "list"}>
                  <ListTool
                    message={props.message}
                    id={props.part.id}
                    tool={props.part.tool}
                    state={props.part.state}
                  />
                </Match>
                <Match when={props.part.tool === "read"}>
                  <ReadTool
                    message={props.message}
                    id={props.part.id}
                    tool={props.part.tool}
                    state={props.part.state}
                  />
                </Match>
                <Match when={props.part.tool === "write"}>
                  <WriteTool
                    message={props.message}
                    id={props.part.id}
                    tool={props.part.tool}
                    state={props.part.state}
                  />
                </Match>
                <Match when={props.part.tool === "edit"}>
                  <EditTool
                    message={props.message}
                    id={props.part.id}
                    tool={props.part.tool}
                    state={props.part.state}
                  />
                </Match>
                <Match when={props.part.tool === "bash"}>
                  <BashTool
                    id={props.part.id}
                    tool={props.part.tool}
                    state={props.part.state}
                    message={props.message}
                  />
                </Match>
                <Match when={props.part.tool === "todowrite"}>
                  <TodoWriteTool
                    message={props.message}
                    id={props.part.id}
                    tool={props.part.tool}
                    state={props.part.state}
                  />
                </Match>
                <Match when={props.part.tool === "webfetch"}>
                  <WebFetchTool
                    message={props.message}
                    id={props.part.id}
                    tool={props.part.tool}
                    state={props.part.state}
                  />
                </Match>
                <Match when={true}>
                  <FallbackTool id={props.part.id} tool={props.part.tool} state={props.part.state} />
                </Match>
              </Switch>
            </div>
          )}
      </div>
    </div>
  )
}

type ToolProps = {
  id: MessageV2.ToolPart["id"]
  tool: MessageV2.ToolPart["tool"]
  state: MessageV2.ToolStateCompleted
  message: MessageV2.Assistant
  isLastPart?: boolean
}

interface Todo {
  id: string
  content: string
  status: "pending" | "in_progress" | "completed"
  priority: "low" | "medium" | "high"
}

function stripWorkingDirectory(filePath?: string, workingDir?: string) {
  if (filePath === undefined || workingDir === undefined) return filePath

  const prefix = workingDir.endsWith("/") ? workingDir : workingDir + "/"

  if (filePath === workingDir) {
    return ""
  }

  if (filePath.startsWith(prefix)) {
    return filePath.slice(prefix.length)
  }

  return filePath
}

function getShikiLang(filename: string) {
  const ext = filename.split(".").pop()?.toLowerCase() ?? ""
  const langs = map.languages(ext)
  const type = langs?.[0]?.toLowerCase()

  const overrides: Record<string, string> = {
    conf: "shellscript",
  }

  return type ? (overrides[type] ?? type) : "plaintext"
}

function getDiagnostics(diagnosticsByFile: Record<string, Diagnostic[]>, currentFile: string): JSX.Element[] {
  const result: JSX.Element[] = []

  if (diagnosticsByFile === undefined || diagnosticsByFile[currentFile] === undefined) return result

  for (const diags of Object.values(diagnosticsByFile)) {
    for (const d of diags) {
      if (d.severity !== 1) continue

      const line = d.range.start.line + 1
      const column = d.range.start.character + 1

      result.push(
        <pre>
          <span data-color="red" data-marker="label">
            Error
          </span>
          <span data-color="dimmed" data-separator>
            [{line}:{column}]
          </span>
          <span>{d.message}</span>
        </pre>,
      )
    }
  }

  return result
}

function formatErrorString(error: string): JSX.Element {
  const errorMarker = "Error: "
  const startsWithError = error.startsWith(errorMarker)

  return startsWithError ? (
    <pre>
      <span data-color="red" data-marker="label" data-separator>
        Error
      </span>
      <span>{error.slice(errorMarker.length)}</span>
    </pre>
  ) : (
    <pre>
      <span data-color="dimmed">{error}</span>
    </pre>
  )
}

export function TodoWriteTool(props: ToolProps) {
  const priority: Record<Todo["status"], number> = {
    in_progress: 0,
    pending: 1,
    completed: 2,
  }
  const todos = createMemo(() =>
    ((props.state.input?.todos ?? []) as Todo[]).slice().sort((a, b) => priority[a.status] - priority[b.status]),
  )
  const starting = () => todos().every((t: Todo) => t.status === "pending")
  const finished = () => todos().every((t: Todo) => t.status === "completed")

  return (
    <>
      <div data-component="tool-title">
        <span data-slot="name">
          <Switch fallback="Updating plan">
            <Match when={starting()}>Creating plan</Match>
            <Match when={finished()}>Completing plan</Match>
          </Switch>
        </span>
      </div>
      <Show when={todos().length > 0}>
        <ul data-component="todos">
          <For each={todos()}>
            {(todo) => (
              <li data-slot="item" data-status={todo.status}>
                <span></span>
                {todo.content}
              </li>
            )}
          </For>
        </ul>
      </Show>
    </>
  )
}

export function GrepTool(props: ToolProps) {
  const count = () => props.state.metadata?.matches
  const pattern = () => props.state.input.pattern

  return (
    <>
      <div data-component="tool-title">
        <span data-slot="name">Grep</span>
        <span data-slot="target">&ldquo;{pattern()}&rdquo;</span>
      </div>
      <div data-component="tool-result">
        <Switch>
          <Match when={count() && count() > 0}>
            <ResultsButton showCopy={count() === 1 ? "1 match" : `${count()} matches`}>
              <ContentText expand compact text={props.state.output} />
            </ResultsButton>
          </Match>
          <Match when={props.state.output}>
            <ContentText expand compact text={props.state.output} data-size="sm" data-color="dimmed" />
          </Match>
        </Switch>
      </div>
    </>
  )
}

export function ListTool(props: ToolProps) {
  const path = createMemo(() =>
    props.state.input?.path !== props.message.path.cwd
      ? stripWorkingDirectory(props.state.input?.path, props.message.path.cwd)
      : props.state.input?.path,
  )

  return (
    <>
      <div data-component="tool-title">
        <span data-slot="name">LS</span>
        <span data-slot="target" title={props.state.input?.path}>
          {path()}
        </span>
      </div>
      <div data-component="tool-result">
        <Switch>
          <Match when={props.state.output}>
            <ResultsButton>
              <ContentText expand compact text={props.state.output} />
            </ResultsButton>
          </Match>
        </Switch>
      </div>
    </>
  )
}

export function WebFetchTool(props: ToolProps) {
  const url = () => props.state.input.url
  const format = () => props.state.input.format
  const hasError = () => props.state.metadata?.error

  return (
    <>
      <div data-component="tool-title">
        <span data-slot="name">Fetch</span>
        <span data-slot="target">{url()}</span>
      </div>
      <div data-component="tool-result">
        <Switch>
          <Match when={hasError()}>
            <div data-component="error">{formatErrorString(props.state.output)}</div>
          </Match>
          <Match when={props.state.output}>
            <ResultsButton>
              <CodeBlock lang={format() || "text"} code={props.state.output} />
            </ResultsButton>
          </Match>
        </Switch>
      </div>
    </>
  )
}

export function ReadTool(props: ToolProps) {
  const filePath = createMemo(() => stripWorkingDirectory(props.state.input?.filePath, props.message.path.cwd))
  const hasError = () => props.state.metadata?.error
  const preview = () => props.state.metadata?.preview

  return (
    <>
      <div data-component="tool-title">
        <span data-slot="name">Read</span>
        <span data-slot="target" title={props.state.input?.filePath}>
          {filePath()}
        </span>
      </div>
      <div data-component="tool-result">
        <Switch>
          <Match when={hasError()}>
            <div data-component="error">{formatErrorString(props.state.output)}</div>
          </Match>
          <Match when={typeof preview() === "string"}>
            <ResultsButton showCopy="Show preview" hideCopy="Hide preview">
              <ContentCode lang={getShikiLang(filePath() || "")} code={preview()} />
            </ResultsButton>
          </Match>
          <Match when={typeof preview() !== "string" && props.state.output}>
            <ResultsButton>
              <ContentText expand compact text={props.state.output} />
            </ResultsButton>
          </Match>
        </Switch>
      </div>
    </>
  )
}

export function WriteTool(props: ToolProps) {
  const filePath = createMemo(() => stripWorkingDirectory(props.state.input?.filePath, props.message.path.cwd))
  const hasError = () => props.state.metadata?.error
  const content = () => props.state.input?.content
  const diagnostics = createMemo(() => getDiagnostics(props.state.metadata?.diagnostics, props.state.input.filePath))

  return (
    <>
      <div data-component="tool-title">
        <span data-slot="name">Write</span>
        <span data-slot="target" title={props.state.input?.filePath}>
          {filePath()}
        </span>
      </div>
      <Show when={diagnostics().length > 0}>
        <div data-component="error">{diagnostics()}</div>
      </Show>
      <div data-component="tool-result">
        <Switch>
          <Match when={hasError()}>
            <div data-component="error">{formatErrorString(props.state.output)}</div>
          </Match>
          <Match when={content()}>
            <ResultsButton showCopy="Show contents" hideCopy="Hide contents">
              <ContentCode lang={getShikiLang(filePath() || "")} code={props.state.input?.content} />
            </ResultsButton>
          </Match>
        </Switch>
      </div>
    </>
  )
}

export function EditTool(props: ToolProps) {
  const diff = () => props.state.metadata?.diff
  const message = () => props.state.metadata?.message
  const hasError = () => props.state.metadata?.error
  const filePath = createMemo(() => stripWorkingDirectory(props.state.input.filePath, props.message.path.cwd))
  const diagnostics = createMemo(() => getDiagnostics(props.state.metadata?.diagnostics, props.state.input.filePath))

  return (
    <>
      <div data-component="tool-title">
        <span data-slot="name">Edit</span>
        <span data-slot="target" title={props.state.input?.filePath}>
          {filePath()}
        </span>
      </div>
      <div data-component="tool-result">
        <Switch>
          <Match when={hasError()}>
            <div data-component="error">{formatErrorString(message() || "")}</div>
          </Match>
          <Match when={diff()}>
            <div data-component="diff">
              <ContentDiff diff={diff()} lang={getShikiLang(filePath() || "")} />
            </div>
          </Match>
        </Switch>
      </div>
      <Show when={diagnostics().length > 0}>
        <div data-component="error">{diagnostics()}</div>
      </Show>
    </>
  )
}

export function BashTool(props: ToolProps) {
  const command = () => props.state.metadata?.title
  const result = () => props.state.metadata?.stdout
  const error = () => props.state.metadata?.stderr

  return (
    <>
      <div data-component="terminal" data-size="sm">
        <div data-slot="body">
          <div data-slot="header">
            <span>{props.state.metadata.description}</span>
          </div>
          <div data-slot="content">
            <ContentCode flush lang="bash" code={props.state.input.command} />
            <ContentCode flush lang="console" code={result() || ""} />
          </div>
        </div>
      </div>
    </>
  )
}

export function GlobTool(props: ToolProps) {
  const count = () => props.state.metadata?.count
  const pattern = () => props.state.input.pattern

  return (
    <>
      <div data-component="tool-title">
        <span data-slot="name">Glob</span>
        <span data-slot="target">&ldquo;{pattern()}&rdquo;</span>
      </div>
      <Switch>
        <Match when={count() && count() > 0}>
          <div data-component="tool-result">
            <ResultsButton showCopy={count() === 1 ? "1 result" : `${count()} results`}>
              <ContentText expand compact text={props.state.output} />
            </ResultsButton>
          </div>
        </Match>
        <Match when={props.state.output}>
          <ContentText expand text={props.state.output} data-size="sm" data-color="dimmed" />
        </Match>
      </Switch>
    </>
  )
}

interface ResultsButtonProps extends ParentProps {
  showCopy?: string
  hideCopy?: string
}
function ResultsButton(props: ResultsButtonProps) {
  const [show, setShow] = createSignal(false)

  return (
    <>
      <button type="button" data-component="button-text" data-more onClick={() => setShow((e) => !e)}>
        <span>{show() ? props.hideCopy || "Hide results" : props.showCopy || "Show results"}</span>
        <span data-slot="icon">
          <Show when={show()} fallback={<IconChevronRight width={11} height={11} />}>
            <IconChevronDown width={11} height={11} />
          </Show>
        </span>
      </button>
      <Show when={show()}>{props.children}</Show>
    </>
  )
}

export function Spacer() {
  return <div data-component="spacer"></div>
}

function Footer(props: ParentProps<{ title: string }>) {
  return (
    <div data-component="content-footer" title={props.title}>
      {props.children}
    </div>
  )
}
