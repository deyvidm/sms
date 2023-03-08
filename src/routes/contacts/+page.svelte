<script lang="ts">
    import ContactList from "$lib/ContactList.svelte";
    import { get50Contacts } from "$lib/pocketbase";
    import { onMount } from "svelte";
    
    onMount(()=>{
        get50Contacts()
    })
</script>

<ContactList/>