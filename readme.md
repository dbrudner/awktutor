# AwkTutor

AwkTutor is an interactive command-line tutorial for learning `awk`, inspired by `vimtutor`. It presents a short rule, command, or concept, shows examples, and allows the user to write an `awk` script to translate given input to the expected output.

## Running AwkTutor

To start the tutorial:

```sh
make run
```

## Lessons

Each lesson consists of:

- A rule/command/concept.
- Examples with explanations.
- A challenge where the user writes an `awk` script to achieve a specific output from a given input.

## Adding a New Lesson

1. Create a new text file in the lessons/ directory.
2. Follow the format of existing lessons for consistency.
3. Update awktutor.go to include the new lesson.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
