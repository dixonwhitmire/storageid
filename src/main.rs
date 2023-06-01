use rand::distributions::{Alphanumeric, DistString};
use regex::Regex;
use std::io;

/// The maximum id length for the generated identifier.
const MAX_ID_LENGTH: usize = 24;

/// create_id generates an id compatible with Azure Storage account naming conventions.
/// An optional prefix may be provided for the id.
fn create_id(prefix: &str, max_id_length: usize) -> Result<String, String> {
    let re = Regex::new(r"\w+").expect("string should be a valid regular expression");
    
    let id_length = max_id_length - prefix.len();
    let id = prefix.to_string() + &Alphanumeric.sample_string(&mut rand::thread_rng(), id_length);

    if re.is_match(&id) {
        return Ok(id)
    } else {
        return Err("Generated id is invalid".to_string());
    }
}

/// Entrypoint for storageid.
/// Accepts prompt from a user for a storageid prefix.
fn main() {
    println!("Please enter a storage account prefix or hit enter for None");
    let mut prefix = String::new();

    io::stdin()
        .read_line(&mut prefix)
        .expect("failed to read line");

    // cleanup the new line character
    prefix = prefix.replace("\n", "");

    match create_id(&prefix, MAX_ID_LENGTH) {
        Ok(value) => println!("New ID = {value}"),
        Err(value) => eprintln!("Error = {value}"),
    }
}
