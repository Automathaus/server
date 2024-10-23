<script lang="ts">
    import Dashboard from '$lib/pages/dashboard.svelte';

    // Wails
    import { StartServer } from '$lib/wailsjs/go/automathaus/AutomathausServer';
    
    // Assets
    import Server from "lucide-svelte/icons/server";
    import Settings2 from "lucide-svelte/icons/settings-2";
    import Info from "lucide-svelte/icons/info";
    import SquareTerminal from "lucide-svelte/icons/square-terminal";
    import LampCeiling from "lucide-svelte/icons/lamp-ceiling";
    import Thermometer from "lucide-svelte/icons/thermometer";
    import Cpu from "lucide-svelte/icons/cpu";
    import Blinds from "lucide-svelte/icons/blinds";

    // Components
    import { Button } from "$lib/components/ui/button/index.js";
    import * as Tooltip from "$lib/components/ui/tooltip/index.js";
    import LogoAutomat from "$lib/components/svg/logoAutomat.svelte";
    import { Toaster } from '$lib/components/ui/sonner';
    import { toast } from "svelte-sonner";
    import * as Card from "$lib/components/ui/card/index.js";
    import DarkmodeButton from '$lib/components/ui/darkmodeButton.svelte';

    async function startServer() {
        console.log('Starting server...');
        try {
            const result = await StartServer();
            console.log(result);
            toast.success(result, {
                description: "Listening for connections...",
                action: {
                    label: "Dismiss",
                    onClick: () => console.info("Undo")
                }
            });
        } catch (err) {
            console.error('Error starting server', err);
            toast.error("Error starting server", {
                description: err as string,
                action: {
                    label: "Dismiss",
                    onClick: () => console.info("Undo")
                }
            });
        }
    }
</script>

<Toaster/>

<div class="grid h-screen w-full pl-[53px]">
    <aside class="fixed left-0 z-20 flex h-full flex-col border-r backdrop-blur-md bg-white/90 dark:bg-zinc-950/50">
        <div class="border-b p-2">
            <Button variant="outline" size="icon" aria-label="Home">
                <LogoAutomat class="size-6"/>
            </Button>
        </div>


        <nav class="grid gap-1 p-2">
            <Tooltip.Root>
                <Tooltip.Trigger asChild let:builder>
                    <Button
                        variant="ghost"
                        size="icon"
                        class="rounded-lg"
                        aria-label="Models"
                        builders={[builder]}
                    >
                        <Cpu class="size-5"/>
                    </Button>
                </Tooltip.Trigger>
                <Tooltip.Content side="right" sideOffset={5}
                    >Nodes</Tooltip.Content
                >
            </Tooltip.Root>

            <Tooltip.Root>
                <Tooltip.Trigger asChild let:builder>
                    <Button
                        variant="ghost"
                        size="icon"
                        class="rounded-lg"
                        aria-label="Models"
                        builders={[builder]}
                    >
                        <SquareTerminal class="size-5"/>
                    </Button>
                </Tooltip.Trigger>
                <Tooltip.Content side="right" sideOffset={5}
                    >Console</Tooltip.Content
                >
            </Tooltip.Root>

            <Tooltip.Root>
                <Tooltip.Trigger asChild let:builder>
                    <Button
                        variant="ghost"
                        size="icon"
                        class="rounded-lg"
                        aria-label="Models"
                        builders={[builder]}
                    >
                        <LampCeiling class="size-5"/>
                    </Button>
                </Tooltip.Trigger>
                <Tooltip.Content side="right" sideOffset={5}
                    >Lights</Tooltip.Content
                >
            </Tooltip.Root>

            <Tooltip.Root>
                <Tooltip.Trigger asChild let:builder>
                    <Button
                        variant="ghost"
                        size="icon"
                        class="mt-auto rounded-lg"
                        aria-label="Account"
                        builders={[builder]}
                    >
                        <Thermometer class="size-5"/>
                    </Button>
                </Tooltip.Trigger>
                <Tooltip.Content side="right" sideOffset={5}
                    >Thermostat</Tooltip.Content
                >
            </Tooltip.Root>

            <Tooltip.Root>
                <Tooltip.Trigger asChild let:builder>
                    <Button
                        variant="ghost"
                        size="icon"
                        class="mt-auto rounded-lg"
                        aria-label="Account"
                        builders={[builder]}
                    >
                        <Blinds class="size-5"/>
                    </Button>
                </Tooltip.Trigger>
                <Tooltip.Content side="right" sideOffset={5}
                    >Roller blinds</Tooltip.Content
                >
            </Tooltip.Root>
        </nav>




        <nav class="mt-auto grid gap-1 p-2">
            <Tooltip.Root>
                <Tooltip.Trigger asChild let:builder>
                    <Button
                        variant="ghost"
                        size="icon"
                        class="mt-auto rounded-lg"
                        aria-label="Help"
                        builders={[builder]}
                    >
                        <Info class="size-5" />
                    </Button>
                </Tooltip.Trigger>
                <Tooltip.Content side="right" sideOffset={5}
                    >Info</Tooltip.Content
                >
            </Tooltip.Root>

            <Tooltip.Root>
                <Tooltip.Trigger asChild let:builder>
                    <Button
                        variant="ghost"
                        size="icon"
                        class="mt-auto rounded-lg"
                        aria-label="Account"
                        builders={[builder]}
                    >
                        <Settings2 class="size-5" />
                    </Button>
                </Tooltip.Trigger>
                <Tooltip.Content side="right" sideOffset={5}
                    >Settings</Tooltip.Content
                >
            </Tooltip.Root>

            <Tooltip.Root>
                <Tooltip.Trigger asChild let:builder>
                    <DarkmodeButton/>
                </Tooltip.Trigger>
                <Tooltip.Content side="right" sideOffset={5}
                    >Settings</Tooltip.Content
                >
            </Tooltip.Root>
        </nav>
    </aside>


    <div class="flex flex-col">
        <header class="backdrop-blur-md bg-white/80 dark:bg-zinc-950/50 sticky top-0 z-10 flex h-[57px] items-center gap-1 border-b pl-5 pr-2">
            <h1 class="text-xl font-semibold">Dashboard</h1>
            <Button size="sm" class="ml-auto gap-1.5 text-sm">
                <Server class="size-3.5"/>
                Start server
            </Button>
        </header>
        <main class="grid flex-1 gap-4 overflow-auto p-4 z-50 grid-rows-2 md:grid-cols-2 lg:grid-cols-3">
            <Dashboard/>
        </main>
    </div>
</div>
