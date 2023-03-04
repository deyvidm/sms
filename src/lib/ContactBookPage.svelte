<script lang="ts">
    import ContactRow from './ContactRow.svelte';
    import {  contacts } from './pocketbase';
    let yesall = true;
</script>

<div class="overflow-x-auto w-full">
    <table class="table w-full">
        <thead>
            <tr>
                <th>
                    <label>
                        <input type="checkbox" class="checkbox" />
                    </label>
                </th>
                <th>Name</th>
                <th>Favorite Color</th>
                <th />
            </tr>
        </thead>

        <tbody>
            {#each $contacts as contact}
                <ContactRow checked={false} first={contact.first_name} last={contact.last_name} id={contact.id} />
            {/each}
        </tbody>
    </table>
</div>
