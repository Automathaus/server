<script>
    import { ModeWatcher } from "mode-watcher";
    import "../app.css";
    import '@fontsource/poppins';
    import CurveThingsScreen from "$lib/components/svg/CurveThingsScreen.svelte";
</script>

<ModeWatcher />
<CurveThingsScreen />

<slot/>
