<script lang="ts">
    import InputBox from './InputBox.svelte';
    import PhoneInputBox from './PhoneInputBox.svelte';
    import { apiClient } from './gin';

    // <input> tags at our disposal
    let firstname: InputBox;
    let lastname: InputBox;
    let phone: PhoneInputBox;
    let submit: HTMLButtonElement;
    let submitLabel: string = 'Create Contact';

    function createContact() {
        if (!verifyInputs()) {
            return;
        }

        apiClient.AddContact(firstname.Get(), lastname.Get(), phone.Get())
            .then((result) => {
                setButtonStatus(result)
            })
            .catch((error) => {
                setButtonStatus(false);
                console.log(error);
                submitLabel = 'shit';
            });
    }

    function setButtonStatus(success: boolean) {
        if (success) {
            submit.classList.add('btn-disabled', 'btn-success');
            submitLabel = 'Success!';
        } else {
            submit.classList.add('btn-disabled');
        }
    }
    function verifyInputs() {
        let results = [
            firstname.ValidateAndMark(),
            lastname.ValidateAndMark(),
            phone.ValidateAndMark(),
        ];
        return results.reduce((a, b) => a && b);
    }
</script>

<div class="form-control w-full max-w-xs">
    <InputBox bind:this={firstname} label="First Name" />
    <InputBox bind:this={lastname} label="Last Name" />
    <PhoneInputBox bind:this={phone} label="Phone" />
    <button bind:this={submit} on:click={createContact} class="mt-12 w-auto btn btn-active"
        >{submitLabel}</button
    >
</div>
