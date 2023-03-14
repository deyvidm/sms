<script lang="ts">
    import ContactList from "$lib/ContactList.svelte";
    import { get50Contacts } from "$lib/pocketbase";
    import { onMount } from "svelte";
    
    // onMount(async ()=>{
    //     get50Contacts()
    // })
    function handleMessage(e){
        console.log(e.detail)
    }
</script>

<ContactList on:message={handleMessage}/>