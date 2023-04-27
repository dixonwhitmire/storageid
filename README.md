# storageid

Generates a valid "ID" for Azure Storage Accounts. A valid Azure Storage Account identifier:

* Is a maximum of 24 characters.
* Consists of lowercase alpha an numeric characters.

## Usage
Use the following commands to test and build the storageid executable. The executable is created in the `target` directory.

```
make test
make build
cd target
```

To run the executable provide a storage account "prefix" and optionally, the desired length of the id. The length defaults to 24 if not provided.

```
% ./storageid intenv
2023/04/27 11:45:23 main: generated storageid intenvd75oh47pnxuqsvdl9u
```

If the prefix violates the Azure naming convention, an error is returned
```
% ./storageid Intenv
2023/04/27 11:46:36 main: an error occurred storageid.New: invalid prefix Intenv
```