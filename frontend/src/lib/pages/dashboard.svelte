<script lang="ts">
    import * as Card from "$lib/components/ui/card/index.js";
</script>

<Card.Root
    class="col-span-2 backdrop-blur-md bg-white/90 dark:bg-zinc-950/50 flex flex-col"
>
    <Card.Header>
        <Card.Title>Temperature and humidity</Card.Title>
        <Card.Description>Over the last 24 hours</Card.Description>
    </Card.Header>
    <Card.Content class="flex-1 min-h-0">
        <div class="h-full">
            <p>Card Content</p>
        </div>
    </Card.Content>
    <Card.Footer>
        <p>Card Footer</p>
    </Card.Footer>
</Card.Root>

<Card.Root
    class="col-span-1 row-span-2 backdrop-blur-md bg-white/90 dark:bg-zinc-950/50 flex flex-col"
>
    <Card.Header>
        <Card.Title>Lights</Card.Title>
        <Card.Description>List of controllable lights</Card.Description>
    </Card.Header>
    <Card.Content class="flex-1 min-h-0">
        <div class="h-full">
            <p>Card Content</p>
        </div>
    </Card.Content>
    <Card.Footer>
        <p>Card Footer</p>
    </Card.Footer>
</Card.Root>

<Card.Root
    class="col-span-1 backdrop-blur-md bg-white/90 dark:bg-zinc-950/50 flex flex-col"
>
    <Card.Header>
        <Card.Title>Temperature</Card.Title>
        <Card.Description>Current temperature</Card.Description>
    </Card.Header>
    <Card.Content class="flex-1 min-h-0">
        <div class="h-full">
            <p>Card Content</p>
        </div>
    </Card.Content>
    <Card.Footer>
        <p>Card Footer</p>
    </Card.Footer>
</Card.Root>

<Card.Root
    class="col-span-1 backdrop-blur-md bg-white/90 dark:bg-zinc-950/50 flex flex-col"
>
    <Card.Header>
        <Card.Title>Humidity</Card.Title>
        <Card.Description>Current humidity</Card.Description>
    </Card.Header>
    <Card.Content class="flex-1 min-h-0">
        <div class="h-full">
            <p>Card Content</p>
        </div>
    </Card.Content>
    <Card.Footer>
        <p>Card Footer</p>
    </Card.Footer>
</Card.Root>
