package catalog

func (c *Control) FindParamById(id string) *Param {
	for i, param := range c.Parameters {
		if param.Id == id {
			return &c.Parameters[i]
		}
	}
	return nil
}

func (c *Control) StatementToMarkdown() string {
	result := ""
	for _, part := range c.Parts {
		if part.Name != "statement" {
			continue
		}
		result += c.partToMarkdown(&part, "")
	}

	return result
}

func (c *Control) partToMarkdown(part *Part, textPrefix string) string {
	result := ""
	if part.Prose != nil {
		label := ""
		for _, prop := range part.Properties {
			if prop.Name == "label" {
				label += prop.Value + " "
				break
			}
		}

		result = textPrefix + label + part.Prose.Raw
		if result[len(result)-1] != '\n' {
			result += "\n"
		}
	}

	prefix := "  " + textPrefix
	if textPrefix == "" {
		prefix = " - "
	}

	for _, child := range part.Parts {
		result += c.partToMarkdown(&child, prefix)
	}

	return result
}
