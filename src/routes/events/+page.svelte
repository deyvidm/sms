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

<h1>Events</h1>
<EventList on:message={handleMessage}/>