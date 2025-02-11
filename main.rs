use doc_comment::module;

// Print all items in the threading module.
for item in module!(std::threading).items() {
    println!("{}", item);
}
