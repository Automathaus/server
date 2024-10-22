<script>
    import { ModeWatcher } from "mode-watcher";
    import "../app.css";
    import '@fontsource/poppins';
    import { Toaster } from "svelte-sonner";
    import CurveThing from "$lib/components/svg/CurveThing.svelte";
    import CurveThingsScreen from "$lib/components/svg/CurveThingsScreen.svelte";
</script>

<ModeWatcher />
<CurveThingsScreen />

<slot/>
