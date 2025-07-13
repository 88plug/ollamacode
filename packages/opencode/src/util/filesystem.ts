import { exists } from "fs/promises"
import { dirname, join, relative } from "path"

export namespace Filesystem {
  export function overlaps(a: string, b: string) {
    const relA = relative(a, b)
    const relB = relative(b, a)
    return !relA || !relA.startsWith("..") || !relB || !relB.startsWith("..")
  }

  export function contains(parent: string, child: string) {
    return relative(parent, child).startsWith("..")
  }

  export async function findUp(target: string, start: string, stop?: string) {
    let current = start
    const result = []
    while (true) {
      const search = join(current, target)
      if (await exists(search)) result.push(search)
      if (stop === current) break
      const parent = dirname(current)
      if (parent === current) break
      current = parent
    }
    return result
  }

  export async function* up(options: { targets: string[]; start: string; stop?: string }) {
    const { targets, start, stop } = options
    let current = start
    while (true) {
      for (const target of targets) {
        const search = join(current, target)
        if (await exists(search)) yield search
      }
      if (stop === current) break
      const parent = dirname(current)
      if (parent === current) break
      current = parent
    }
  }

  export async function globUp(pattern: string, start: string, stop?: string) {
    let current = start
    const result = []
    while (true) {
      try {
        const glob = new Bun.Glob(pattern)
        for await (const match of glob.scan({
          cwd: current,
          onlyFiles: true,
          dot: true,
        })) {
          result.push(join(current, match))
        }
      } catch {
        // Skip invalid glob patterns
      }
      if (stop === current) break
      const parent = dirname(current)
      if (parent === current) break
      current = parent
    }
    return result
  }

  export async function findDown(target: string, start: string, maxDepth: number = 3) {
    const result = []
    const glob = new Bun.Glob(`**/${target}`)
    try {
      for await (const match of glob.scan({
        cwd: start,
        onlyFiles: true,
        dot: true,
      })) {
        const fullPath = join(start, match)
        const depth = match.split('/').length - 1
        if (depth <= maxDepth) {
          result.push(fullPath)
        }
      }
    } catch {
      // Skip if glob fails
    }
    return result
  }
}
