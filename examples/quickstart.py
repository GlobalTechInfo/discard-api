from discardapi import DiscardAPI

client = DiscardAPI("YOUR_API_KEY")
result = client.markdown_to_html("# Hello DiscardAPI")

print(result)
