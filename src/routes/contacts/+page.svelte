<script lang="ts">
    import ContactList from "$lib/ContactList.svelte";
    import { get50Contacts } from "$lib/pocketbase";
    import { onMount } from "svelte";
    

    function handleMessage(e){
    }
</script>

<ContactList on:message={handleMessage}/>