<script lang="ts">
    import { ArrowRight } from 'lucide-svelte';
    import { Server } from 'lucide-svelte';
    import { toast } from "svelte-sonner";

    import { StartServer } from '$lib/wailsjs/go/main/AutomathausServer';
    import Button from "$lib/components/ui/button/button.svelte";
    import { Toaster } from '$lib/components/ui/sonner';

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

<div class="w-full h-svh flex justify-center items-center">
    <Button on:click={startServer} class="hover:scale-110 ease-in-out transition-transform duration-300">
        <Server class="mr-2 h-5 w-5"/>
        Start server
        <ArrowRight class="ml-2 h-5 w-5"/>
    </Button>
</div>
