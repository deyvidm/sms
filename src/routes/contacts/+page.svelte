<script>
    import { Record } from 'pocketbase';
    import ContactList from '$lib/ContactList.svelte';

    /** @type {import('./$types').PageData} */
    export let data;

    function handleMessage(event) {
        let recipients = event.detail.recipients;
    }


</script>
<ContactList contacts={data.contacts} on:message={handleMessage} />
