<script lang="ts">
    import ContactList from "$lib/ContactList.svelte";
    

    function handleMessage(e){
    }
</script>

<ContactList on:message={handleMessage}/>