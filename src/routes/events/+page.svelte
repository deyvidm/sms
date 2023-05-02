<script lang="ts">
    import { get50Events } from "$lib/pocketbase";
    import { onMount } from "svelte";
    import EventList from "$lib/EventList.svelte";
    
    onMount(()=>{
        get50Events()
    })
    function handleMessage(e){
    }
</script>

<EventList on:message={handleMessage}/>