<script lang="ts">
    import { Menu } from 'lucide-svelte';
    import { SunMoon } from 'lucide-svelte';

    import Github from "lucide-svelte/icons/github";
    import Newspaper from "lucide-svelte/icons/newspaper";
    import BookText from "lucide-svelte/icons/book-text";
    import House from "lucide-svelte/icons/house";

    import DarkmodeButton from "../darkmodeButton.svelte";
    import LogoAutomat from "../../svg/logoAutomat.svelte";

    import * as DropdownMenu from "$lib/components/ui/dropdown-menu/index.js";
    import { Button } from "$lib/components/ui/button/index.js";

    import { toggleMode } from "mode-watcher";
</script>

<nav class="fixed p-3 md:p-5 w-full flex items-center justify-between backdrop-blur-md bg-white/80 dark:bg-zinc-950/80 z-50">
    <div class="flex items-center space-x-3 ">
         <LogoAutomat class="size-7 lg:size-9"/>
        <h2 class="text-xl lg:text-2xl">
            Automathaus<span class="text-zinc-500/30">.dev</span>
        </h2>
    </div>

    
    <div class="md:flex items-center space-x-3">
        <Button variant="outline" class="hidden md:flex">
            <House class="mr-2 h-4 w-4"/>
            Home
        </Button>
        
        <Button class="hidden md:flex">
            <Github class="mr-2 h-4 w-4"/>
            Github
        </Button>

        <DarkmodeButton class="hidden md:flex" />

        <DropdownMenu.Root>
            <DropdownMenu.Trigger asChild let:builder>
                <Button builders={[builder]} variant="ghost" size="icon" class="md:hidden">
                    <Menu class="size-7" />
                    <span class="sr-only">Toggle theme</span>
                </Button>
            </DropdownMenu.Trigger>
            <DropdownMenu.Content align="end">
                <DropdownMenu.Item >
                    <House class="mr-2 h-4 w-4"/>
                    <span>Home</span>
                </DropdownMenu.Item>
                
                <DropdownMenu.Item >
                    <BookText class="mr-2 h-4 w-4"/>
                    <span>Docs</span>
                </DropdownMenu.Item>
    
                <DropdownMenu.Item >
                    <Newspaper class="mr-2 h-4 w-4"/>
                    <span>News</span>
                </DropdownMenu.Item>
    
                <DropdownMenu.Item >
                    <Github class="mr-2 h-4 w-4"/>
                    <span>Github</span>
                </DropdownMenu.Item>

                <DropdownMenu.Separator />

                <DropdownMenu.Item on:click={toggleMode}>
                    <SunMoon class="mr-2 h-4 w-4"/>
                    <span>Toggle theme</span>
                </DropdownMenu.Item>
            </DropdownMenu.Content>
        </DropdownMenu.Root>
    </div>
</nav>